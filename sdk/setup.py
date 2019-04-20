import setuptools

setuptools.setup(
    name="kubeflow-metadata",
    version="0.1",
    packages=setuptools.find_packages(),

    install_requires=['ml_metadata>=0.13.2'],

    author="Kubeflow",
    description="Prototype KFP Metadata SDK",
)