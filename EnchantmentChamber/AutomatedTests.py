# Automated Tests

def cast_fireball():
    return "Fireball cast successfully"

def cast_heal():
    return "Heal spell invoked"

def cast_teleport():
    return "Teleportation successful"

import unittest

class TestSpellFunctions(unittest.TestCase):
    def test_fireball_spell(self):
        # Test fireball spell
        result = cast_fireball()
        self.assertEqual(result, "Fireball cast successfully")

    def test_heal_spell(self):
        # Test heal spell
        result = cast_heal()
        self.assertEqual(result, "Heal spell invoked")

    def test_teleport_spell(self):
        # Test teleport spell
        result = cast_teleport()
        self.assertEqual(result, "Teleportation successful")

if __name__ == '__main__':
    unittest.main()
