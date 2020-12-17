package com.atlas.eso.database.transformer;

import com.atlas.eso.builder.EquipmentDataBuilder;
import com.atlas.eso.entity.Equipment;
import com.atlas.eso.model.EquipmentData;

import transformer.SqlTransformer;

public class EquipmentDataTransformer implements SqlTransformer<EquipmentData, Equipment> {
   @Override
   public EquipmentData transform(Equipment equipment) {
      return new EquipmentDataBuilder()
            .setId(equipment.getId())
            .setItemId(equipment.getItemId())
            .setStrength(equipment.getStrength())
            .setDexterity(equipment.getDexterity())
            .setIntelligence(equipment.getIntelligence())
            .setLuck(equipment.getLuck())
            .setHp(equipment.getHp())
            .setMp(equipment.getMp())
            .setWeaponAttack(equipment.getWeaponAttack())
            .setWeaponDefense(equipment.getWeaponDefense())
            .setMagicAttack(equipment.getMagicAttack())
            .setMagicDefense(equipment.getMagicDefense())
            .setAccuracy(equipment.getAccuracy())
            .setAvoidability(equipment.getAvoidability())
            .setHands(equipment.getHands())
            .setSpeed(equipment.getSpeed())
            .setJump(equipment.getJump())
            .setSlots(equipment.getSlots())
            .build();
   }
}
