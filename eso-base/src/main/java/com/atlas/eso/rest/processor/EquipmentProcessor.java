package com.atlas.eso.rest.processor;

import javax.ws.rs.core.Response;

import com.app.rest.util.stream.Mappers;
import com.atlas.eso.attribute.EquipmentAttributes;

import builder.ResultBuilder;

public final class EquipmentProcessor {
   private EquipmentProcessor() {
   }

   public static ResultBuilder create(EquipmentAttributes attributes) {
      return com.atlas.eso.processor.EquipmentProcessor.create(attributes.itemId(), attributes.strength(), attributes.dexterity(),
            attributes.intelligence(), attributes.luck(), attributes.weaponAttack(), attributes.weaponDefense(),
            attributes.magicAttack(), attributes.magicDefense(), attributes.accuracy(), attributes.avoidability(),
            attributes.speed(), attributes.jump(), attributes.hp(), attributes.mp(), attributes.slots())
            .map(ResultObjectFactory::create)
            .map(Mappers::singleCreatedResult)
            .orElse(new ResultBuilder(Response.Status.FORBIDDEN));
   }

   public static ResultBuilder get(int id) {
      return com.atlas.eso.processor.EquipmentProcessor.get(id)
            .map(ResultObjectFactory::create)
            .map(Mappers::singleOkResult)
            .orElse(new ResultBuilder(Response.Status.FORBIDDEN));
   }
}
