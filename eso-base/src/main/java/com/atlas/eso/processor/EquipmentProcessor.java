package com.atlas.eso.processor;

import java.util.Optional;
import java.util.function.Consumer;

import com.atlas.eso.builder.EquipmentDataBuilder;
import com.atlas.eso.database.administrator.EquipmentAdministrator;
import com.atlas.eso.database.provider.EquipmentProvider;
import com.atlas.eso.model.EquipmentData;
import com.atlas.iis.attribute.EquipmentAttributes;
import com.atlas.iis.constant.RestConstants;
import com.atlas.shared.rest.UriBuilder;

import database.Connection;
import rest.DataContainer;

public final class EquipmentProcessor {
   private EquipmentProcessor() {
   }

   public static Optional<EquipmentData> create(int itemId, Integer strength, Integer dexterity, Integer intelligence,
                                                Integer luck, Integer weaponAttack, Integer weaponDefense, Integer magicAttack,
                                                Integer magicDefense, Integer accuracy, Integer avoidability, Integer speed,
                                                Integer jump, Integer hp, Integer mp, Integer slots) {
      EquipmentDataBuilder equipmentBuilder = new EquipmentDataBuilder().setItemId(itemId);

      UriBuilder.service(RestConstants.SERVICE)
            .pathParam("equipment", itemId)
            .getRestClient(EquipmentAttributes.class)
            .getWithResponse()
            .result()
            .flatMap(DataContainer::data)
            .ifPresent(body -> {
               equipmentBuilder.setStrength(body.getAttributes().strength());
               equipmentBuilder.setDexterity(body.getAttributes().dexterity());
               equipmentBuilder.setIntelligence(body.getAttributes().intelligence());
               equipmentBuilder.setLuck(body.getAttributes().luck());
               equipmentBuilder.setWeaponAttack(body.getAttributes().weaponAttack());
               equipmentBuilder.setWeaponDefense(body.getAttributes().weaponDefense());
               equipmentBuilder.setMagicAttack(body.getAttributes().magicAttack());
               equipmentBuilder.setMagicDefense(body.getAttributes().magicDefense());
               equipmentBuilder.setAccuracy(body.getAttributes().accuracy());
               equipmentBuilder.setAvoidability(body.getAttributes().avoidability());
               equipmentBuilder.setSpeed(body.getAttributes().speed());
               equipmentBuilder.setJump(body.getAttributes().jump());
               equipmentBuilder.setHp(body.getAttributes().hp());
               equipmentBuilder.setMp(body.getAttributes().mp());
               equipmentBuilder.setSlots(body.getAttributes().slots());
            });

      setIfNotNull(equipmentBuilder::setStrength, strength);
      setIfNotNull(equipmentBuilder::setDexterity, dexterity);
      setIfNotNull(equipmentBuilder::setIntelligence, intelligence);
      setIfNotNull(equipmentBuilder::setLuck, luck);
      setIfNotNull(equipmentBuilder::setWeaponAttack, weaponAttack);
      setIfNotNull(equipmentBuilder::setWeaponDefense, weaponDefense);
      setIfNotNull(equipmentBuilder::setMagicAttack, magicAttack);
      setIfNotNull(equipmentBuilder::setMagicDefense, magicDefense);
      setIfNotNull(equipmentBuilder::setAccuracy, accuracy);
      setIfNotNull(equipmentBuilder::setAvoidability, avoidability);
      setIfNotNull(equipmentBuilder::setSpeed, speed);
      setIfNotNull(equipmentBuilder::setJump, jump);
      setIfNotNull(equipmentBuilder::setHp, hp);
      setIfNotNull(equipmentBuilder::setMp, mp);
      setIfNotNull(equipmentBuilder::setSlots, slots);

      EquipmentData equipment = equipmentBuilder.build();

      return Connection.instance()
            .element(entityManager -> EquipmentAdministrator.create(entityManager, equipment));
   }

   protected static <T> void setIfNotNull(Consumer<T> setter, T value) {
      if (value != null) {
         setter.accept(value);
      }
   }

   public static Optional<EquipmentData> get(int id) {
      return Connection.instance().element(entityManager -> EquipmentProvider.get(entityManager, id));
   }
}
