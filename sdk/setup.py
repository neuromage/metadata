import setuptools

setuptools.setup(
    name="kubeflow-metadata",
    version="0.1",
    packages=setuptools.find_packages(),

    # Project uses reStructuredText, so ensure that the docutils get
    # installed or upgraded on the target machine
    install_requires=['ml_metadata>=0.13.2'],

    # metadata to display on PyPI
    author="Kubeflow",
    description="Prototype KFP Metadata SDK",
)