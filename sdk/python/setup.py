from setuptools import setup, find_packages

setup(
    name="kfmd",
    version="0.1.4",
    packages=find_packages(),
    install_requires=['cloudpickle'],
    author="Kubeflow",
    description="Prototype KFP Metadata SDK",
)