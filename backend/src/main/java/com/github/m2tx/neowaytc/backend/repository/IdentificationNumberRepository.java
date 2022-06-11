package com.github.m2tx.neowaytc.backend.repository;

import java.util.UUID;

import org.springframework.data.jpa.repository.JpaRepository;

import com.github.m2tx.neowaytc.backend.model.IdentificationNumber;

public interface IdentificationNumberRepository extends JpaRepository<IdentificationNumber, UUID> {  

}
