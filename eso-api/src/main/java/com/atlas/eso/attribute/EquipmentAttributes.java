package com.atlas.eso.attribute;

import rest.AttributeResult;

public record EquipmentAttributes(Integer itemId, Integer strength, Integer dexterity, Integer intelligence, Integer luck,
                                  Integer hp, Integer mp, Integer weaponAttack, Integer magicAttack, Integer weaponDefense,
                                  Integer magicDefense, Integer accuracy, Integer avoidability, Integer hands, Integer speed,
                                  Integer jump, Integer slots) implements AttributeResult {
}
