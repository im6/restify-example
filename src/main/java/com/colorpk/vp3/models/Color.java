package com.colorpk.vp3.models;

import java.sql.Timestamp;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.Id;
import javax.persistence.Table;
import javax.persistence.Column;

@Entity
@Table(name = "colorpk_color")
public class Color {
  @Id
  @GeneratedValue
  @Column(name = "id")
  private Integer id;

  @Column(name = "like")
  private Integer like;

  @Column(name = "color")
  private String color;

  @Column(name = "userid")
  private Integer userId;

  @Column(name = "username")
  private String username;

  @Column(name = "createdate")
  private Timestamp createDate;

  @Column(name = "display")
  private Boolean display;

  public Color() {
  }

  public Integer getId() {
    return id;
  }

  public Integer getLike() {
    return like;
  }

  public String getColor() {
    return color;
  }

  public Integer getUserId() {
    return userId;
  }

  public String getUsername() {
    return username;
  }

  public Timestamp getCreateDate() {
    return createDate;
  }

  public Boolean getDisplay() {
    return display;
  }

}