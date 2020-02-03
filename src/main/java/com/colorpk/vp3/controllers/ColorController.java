package com.colorpk.vp3.controllers;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;

import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.SQLException;

@RestController
public class ColorController {
  @GetMapping("/color/{id}")
  public String getOneColor(@PathVariable(name = "id", required = true) String id) {
    System.out.println("hello world");
    Connection conn = null;
    try {
      String connStr = String.format("jdbc:mysql://%s/%s?user=%s&password=%s", System.getenv("SQL_HOST"),
          System.getenv("SQL_DATABASE"), System.getenv("SQL_USERNAME"), System.getenv("SQL_PASSWORD"));
      // System.out.println(connStr);
      conn = DriverManager.getConnection(connStr);

      // Do something with the Connection
    } catch (SQLException ex) {
      // handle any errors
      System.out.println("SQLException: " + ex.getMessage());
      System.out.println("SQLState: " + ex.getSQLState());
      System.out.println("VendorError: " + ex.getErrorCode());
    }
    return "greeting " + id;
  }
}