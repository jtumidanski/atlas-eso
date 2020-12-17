package com.atlas.eso.database.provider;

import java.util.Optional;
import javax.persistence.EntityManager;
import javax.persistence.TypedQuery;

import com.app.database.util.QueryProviderUtil;
import com.atlas.eso.database.transformer.EquipmentDataTransformer;
import com.atlas.eso.entity.Equipment;
import com.atlas.eso.model.EquipmentData;

public final class EquipmentProvider {
   private EquipmentProvider() {
   }

   public static Optional<EquipmentData> get(EntityManager entityManager, int id) {
      TypedQuery<Equipment> query = entityManager.createQuery("SELECT e FROM Equipment e WHERE e.id = :id", Equipment.class);
      query.setParameter("id", id);
      return QueryProviderUtil.optionalElement(query, new EquipmentDataTransformer());
   }
}
