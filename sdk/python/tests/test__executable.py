from kfmd import executable, default_config
import unittest

def add(a, b, outputs):
    outputs['sum'] = a + b

counter = 0
def stateful_increment(outputs):
    global counter
    counter += 1
    outputs['val'] = counter

class ExecutableTest(unittest.TestCase):

    def tearDown(self):
        default_config.artifact_store.clean()
        
    def test_executable_calls_fn(self):
        add_op = executable(add)(a = 1, b = 2)

        self.assertEqual(3, add_op.outputs['sum'])

    def test_executable_cache_hit(self):
        global counter
        increment_op1 = executable(stateful_increment)()
        # Cache hit
        increment_op2 = executable(stateful_increment)()

        self.assertEqual(1, increment_op1.outputs['val'])
        self.assertEqual(1, increment_op2.outputs['val'])
        self.assertEqual(1, counter)
        stateful_increment({})
        self.assertEqual(2, counter)