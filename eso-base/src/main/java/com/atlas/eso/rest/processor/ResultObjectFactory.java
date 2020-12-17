package com.atlas.eso.rest.processor;

import com.atlas.eso.attribute.EquipmentAttributes;
import com.atlas.eso.builder.EquipmentAttributesBuilder;
import com.atlas.eso.model.EquipmentData;

import builder.ResultObjectBuilder;

public final class ResultObjectFactory {
   private ResultObjectFactory() {
   }

   public static ResultObjectBuilder create(EquipmentData data) {
      return new ResultObjectBuilder(EquipmentAttributes.class, data.id())
            .setAttribute(new EquipmentAttributesBuilder()
                  .setItemId(data.itemId())
                  .setStrength(data.strength())
                  .setDexterity(data.dexterity())
                  .setIntelligence(data.intelligence())
                  .setLuck(data.luck())
                  .setHp(data.hp())
                  .setMp(data.mp())
                  .setWeaponAttack(data.weaponAttack())
                  .setMagicAttack(data.magicAttack())
                  .setWeaponDefense(data.weaponDefense())
                  .setMagicDefense(data.magicDefense())
                  .setAccuracy(data.accuracy())
                  .setAvoidability(data.avoidability())
                  .setHands(data.hands())
                  .setSpeed(data.speed())
                  .setJump(data.jump())
                  .setSlots(data.slots())
            );
   }
}
