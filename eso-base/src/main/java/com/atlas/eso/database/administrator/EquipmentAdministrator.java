package com.atlas.eso.database.administrator;

import java.util.Optional;
import javax.persistence.EntityManager;

import com.app.database.util.QueryAdministratorUtil;
import com.atlas.eso.database.transformer.EquipmentDataTransformer;
import com.atlas.eso.entity.Equipment;
import com.atlas.eso.model.EquipmentData;

public class EquipmentAdministrator {
   private EquipmentAdministrator() {
   }

   public static Optional<EquipmentData> create(EntityManager entityManager, EquipmentData equipmentData) {
      Equipment equipment = new Equipment();
      equipment.setItemId(equipmentData.itemId());
      equipment.setStrength(equipmentData.strength());
      equipment.setDexterity(equipmentData.dexterity());
      equipment.setIntelligence(equipmentData.intelligence());
      equipment.setLuck(equipmentData.luck());
      equipment.setHp(equipmentData.hp());
      equipment.setMp(equipmentData.mp());
      equipment.setWeaponAttack(equipmentData.weaponAttack());
      equipment.setWeaponDefense(equipmentData.weaponDefense());
      equipment.setMagicAttack(equipmentData.magicAttack());
      equipment.setMagicDefense(equipmentData.magicDefense());
      equipment.setAccuracy(equipmentData.accuracy());
      equipment.setAvoidability(equipmentData.avoidability());
      equipment.setHands(equipmentData.hands());
      equipment.setSpeed(equipmentData.speed());
      equipment.setJump(equipmentData.jump());
      equipment.setSlots(equipmentData.slots());
      return Optional.of(QueryAdministratorUtil.insertAndReturn(entityManager, equipment, new EquipmentDataTransformer()));
   }
}
