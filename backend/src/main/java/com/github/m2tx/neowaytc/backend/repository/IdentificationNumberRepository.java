package com.github.m2tx.neowaytc.backend.repository;

import java.util.UUID;

import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.domain.Specification;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.JpaSpecificationExecutor;

import com.github.m2tx.neowaytc.backend.model.IdentificationNumber;

public interface IdentificationNumberRepository extends JpaRepository<IdentificationNumber, UUID>, JpaSpecificationExecutor<IdentificationNumber>  {

    Page<IdentificationNumber> findAll(Specification<IdentificationNumber> spec, Pageable pageable);

}
