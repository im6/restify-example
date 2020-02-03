package com.colorpk.vp3.models;

import java.util.Date;

public class Color {
  public final int id;
  public final int like;
  public final String color;
  public final String userId;
  public final String username;
  public final Date createdDate;

  public Boolean display;

  public Color(int id, int like, String color, String userId, String username, Date createdDate) {
    this.id = id;
    this.like = like;
    this.color = color;
    this.userId = userId;
    this.username = username;
    this.createdDate = createdDate;
  }
}