package com.colorpk.vp3.repository;

import com.colorpk.vp3.models.Color;
import org.springframework.data.jpa.repository.JpaRepository;

public interface ColorsRepository extends JpaRepository<Color, Integer> {
}