import kfmd

def write_numpy_array(path, narray):
    import numpy
    import os
    os.makedirs(os.path.dirname(path), exist_ok=True)
    numpy.save(path, narray)

def download_data(outputs):
    import tensorflow as tf
    (x_train, y_train), (x_test, y_test) = tf.keras.datasets.mnist.load_data()
    write_numpy_array(outputs.artifact('x_train.npy')['uri'], x_train)
    write_numpy_array(outputs.artifact('y_train.npy')['uri'], y_train)
    write_numpy_array(outputs.artifact('x_test.npy')['uri'], x_test)
    write_numpy_array(outputs.artifact('y_test.npy')['uri'], y_test)

def reshape(x_train, x_test, outputs):
    import numpy
    import json
    x_train = numpy.load(x_train['uri'])
    x_test = numpy.load(x_test['uri'])

    # Reshaping the array to 4-dims so that it can work with the Keras API
    x_train = x_train.reshape(x_train.shape[0], 28, 28, 1)
    x_test = x_test.reshape(x_test.shape[0], 28, 28, 1)
    input_shape = (28, 28, 1)
    # Making sure that the values are float so that we can get decimal points after division
    x_train = x_train.astype('float32')
    x_test = x_test.astype('float32')
    # Normalizing the RGB codes by dividing it to the max RGB value.
    x_train /= 255
    x_test /= 255
    print('x_train shape:', x_train.shape)
    print('Number of images in x_train', x_train.shape[0])
    print('Number of images in x_test', x_test.shape[0])

    write_numpy_array(outputs.artifact('x_train.npy')['uri'], x_train)
    write_numpy_array(outputs.artifact('x_test.npy')['uri'], x_test)
    outputs['input_shape'] = json.dumps(input_shape)

def prepare_model(input_shape, outputs):
    import tensorflow as tf
    import json
    import os
    input_shape = json.loads(input_shape)

    # Importing the required Keras modules containing model and layers
    from keras.models import Sequential
    from keras.layers import Dense, Conv2D, Dropout, Flatten, MaxPooling2D
    # Creating a Sequential Model and adding the layers
    model = Sequential()
    model.add(Conv2D(28, kernel_size=(3,3), input_shape=input_shape))
    model.add(MaxPooling2D(pool_size=(2, 2)))
    model.add(Flatten()) # Flattening the 2D arrays for fully connected layers
    model.add(Dense(128, activation=tf.nn.relu))
    model.add(Dropout(0.2))
    model.add(Dense(10,activation=tf.nn.softmax))

    model_path = outputs.artifact('model.json')['uri']
    tf.gfile.MakeDirs(os.path.dirname(model_path))
    with tf.gfile.Open(model_path, 'w') as f:
        f.write(model.to_json())

def train_model(untrained_model, x_train, y_train, epochs, outputs):
    import tensorflow as tf
    from keras.models import model_from_json
    import numpy
    import os

    x_train = numpy.load(x_train['uri'])
    y_train = numpy.load(y_train['uri'])
    model = model_from_json(tf.gfile.GFile(untrained_model['uri']).read())

    model.compile(optimizer='adam', 
              loss='sparse_categorical_crossentropy', 
              metrics=['accuracy'])
    model.fit(x=x_train,y=y_train, epochs=epochs)

    model_path = outputs.artifact('model.h5')['uri']
    tf.gfile.MakeDirs(os.path.dirname(model_path))
    model.save(model_path)

def evaluate_model(trained_model, x_test, y_test, outputs):
    import tensorflow as tf
    from keras.models import load_model
    import numpy

    x_test = numpy.load(x_test['uri'])
    y_test = numpy.load(y_test['uri'])
    model = load_model(trained_model['uri'])

    scores = model.evaluate(x_test, y_test)

    model = outputs.artifact('model.h5')
    model['uri'] = trained_model['uri']
    model['loss'] = scores[0]
    model['accuracy'] = scores[1]

def main():
    kfmd.config(executor = kfmd.executors.DockerExecutor(
        base_image = 'gcr.io/deeplearning-platform-release/tf-cpu.1-13:latest'))
    download_op = kfmd.executable(download_data)()
    reshape_op = kfmd.executable(reshape)(
        x_train = download_op.outputs['x_train.npy'],
        x_test = download_op.outputs['x_test.npy']
    )
    prepare_model_op = kfmd.executable(prepare_model)(
        input_shape = reshape_op.outputs['input_shape'])
    train_model_op = kfmd.executable(train_model)(
        untrained_model = prepare_model_op.outputs['model.json'], 
        x_train = reshape_op.outputs['x_train.npy'],
        y_train = download_op.outputs['y_train.npy'],
        epochs = 1)
    evaluate_model_op = kfmd.executable(evaluate_model)(
        trained_model = train_model_op.outputs['model.h5'],
        x_test = reshape_op.outputs['x_test.npy'],
        y_test = download_op.outputs['y_test.npy']
    )
    print(evaluate_model_op.outputs['model.h5'])

if __name__ == "__main__":
    main()