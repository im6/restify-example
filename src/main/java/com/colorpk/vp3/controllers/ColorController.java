package com.colorpk.vp3.controllers;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.beans.factory.annotation.Autowired;

import com.colorpk.vp3.models.Color;
import com.colorpk.vp3.repository.ColorsRepository;

import java.util.ArrayList;
import java.util.List;

@RestController
public class ColorController {
  @Autowired
  ColorsRepository colorRepository;

  @GetMapping("/colors")
  public List<Color> getAllColors() {
    return colorRepository.findAll();
  }

  @GetMapping("/color/{id}")
  public List<Color> getOneColor(@PathVariable(name = "id", required = true) Integer id) {
    System.out.println(id);
    ArrayList<Integer> selectedIds = new ArrayList<Integer>() {
      {
        add(id);
      }
    };
    return colorRepository.findAllById(selectedIds);
  }
}