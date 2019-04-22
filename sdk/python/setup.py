from setuptools import setup, find_packages

REQUIRES = []
with open('requirements.txt') as f:
    REQUIRES = f.readlines()

with open('test-requirements.txt') as f:
    TESTS_REQUIRES = f.readlines()

setup(
    name="kubeflow-metadata",
    version="0.1",
    packages=find_packages(),
    install_requires=REQUIRES,
    tests_require=TESTS_REQUIRES,
    author="Kubeflow",
    description="Prototype KFP Metadata SDK",
)