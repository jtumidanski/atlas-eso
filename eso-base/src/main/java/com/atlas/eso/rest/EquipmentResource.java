package com.atlas.eso.rest;

import javax.ws.rs.Consumes;
import javax.ws.rs.POST;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;

import com.atlas.eso.attribute.EquipmentAttributes;
import com.atlas.eso.rest.processor.EquipmentProcessor;

import rest.InputBody;

@Path("equipment")
public class EquipmentResource {
   @POST
   @Path("")
   @Consumes(MediaType.APPLICATION_JSON)
   @Produces(MediaType.APPLICATION_JSON)
   public Response create(InputBody<EquipmentAttributes> inputBody) {
      return EquipmentProcessor.create(inputBody.attributes()).build();
   }
}
