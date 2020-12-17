package com.atlas.eso.model;

public record EquipmentData(int id, int itemId, int strength, int dexterity, int intelligence, int luck, int hp, int mp,
                            int weaponAttack, int magicAttack, int weaponDefense, int magicDefense, int accuracy, int avoidability,
                            int hands, int speed, int jump, int slots) {
}
