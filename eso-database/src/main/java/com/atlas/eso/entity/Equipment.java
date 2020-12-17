package com.atlas.eso.entity;

import java.io.Serializable;
import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;

@Entity
public class Equipment implements Serializable {
   private static final long serialVersionUID = 1L;

   @Id
   @GeneratedValue(strategy = GenerationType.IDENTITY)
   private Integer id;

   @Column(nullable = false)
   private Integer itemId;

   @Column(nullable = false)
   private Integer strength;

   @Column(nullable = false)
   private Integer dexterity;

   @Column(nullable = false)
   private Integer intelligence;

   @Column(nullable = false)
   private Integer luck;

   @Column(nullable = false)
   private Integer hp;

   @Column(nullable = false)
   private Integer mp;

   @Column(nullable = false)
   private Integer weaponAttack;

   @Column(nullable = false)
   private Integer magicAttack;

   @Column(nullable = false)
   private Integer weaponDefense;

   @Column(nullable = false)
   private Integer magicDefense;

   @Column(nullable = false)
   private Integer accuracy;

   @Column(nullable = false)
   private Integer avoidability;

   @Column(nullable = false)
   private Integer hands;

   @Column(nullable = false)
   private Integer speed;

   @Column(nullable = false)
   private Integer jump;

   @Column(nullable = false)
   private Integer slots;

   public Integer getId() {
      return id;
   }

   public Integer getItemId() {
      return itemId;
   }

   public void setItemId(Integer itemId) {
      this.itemId = itemId;
   }

   public Integer getStrength() {
      return strength;
   }

   public void setStrength(Integer strength) {
      this.strength = strength;
   }

   public Integer getDexterity() {
      return dexterity;
   }

   public void setDexterity(Integer dexterity) {
      this.dexterity = dexterity;
   }

   public Integer getIntelligence() {
      return intelligence;
   }

   public void setIntelligence(Integer intelligence) {
      this.intelligence = intelligence;
   }

   public Integer getLuck() {
      return luck;
   }

   public void setLuck(Integer luck) {
      this.luck = luck;
   }

   public Integer getHp() {
      return hp;
   }

   public void setHp(Integer hp) {
      this.hp = hp;
   }

   public Integer getMp() {
      return mp;
   }

   public void setMp(Integer mp) {
      this.mp = mp;
   }

   public Integer getWeaponAttack() {
      return weaponAttack;
   }

   public void setWeaponAttack(Integer weaponAttack) {
      this.weaponAttack = weaponAttack;
   }

   public Integer getMagicAttack() {
      return magicAttack;
   }

   public void setMagicAttack(Integer magicAttack) {
      this.magicAttack = magicAttack;
   }

   public Integer getWeaponDefense() {
      return weaponDefense;
   }

   public void setWeaponDefense(Integer weaponDefense) {
      this.weaponDefense = weaponDefense;
   }

   public Integer getMagicDefense() {
      return magicDefense;
   }

   public void setMagicDefense(Integer magicDefense) {
      this.magicDefense = magicDefense;
   }

   public Integer getAccuracy() {
      return accuracy;
   }

   public void setAccuracy(Integer accuracy) {
      this.accuracy = accuracy;
   }

   public Integer getAvoidability() {
      return avoidability;
   }

   public void setAvoidability(Integer avoidability) {
      this.avoidability = avoidability;
   }

   public Integer getHands() {
      return hands;
   }

   public void setHands(Integer hands) {
      this.hands = hands;
   }

   public Integer getSpeed() {
      return speed;
   }

   public void setSpeed(Integer speed) {
      this.speed = speed;
   }

   public Integer getJump() {
      return jump;
   }

   public void setJump(Integer jump) {
      this.jump = jump;
   }

   public Integer getSlots() {
      return slots;
   }

   public void setSlots(Integer slots) {
      this.slots = slots;
   }
}
