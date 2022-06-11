package com.github.m2tx.neowaytc.backend.model;

import java.util.UUID;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.Table;

import com.github.m2tx.neowaytc.backend.validator.CpfCnpjValidation;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Entity
@Table(name = "identification_numbers")
@Data
@NoArgsConstructor
@AllArgsConstructor
@Builder
public class IdentificationNumber {

    @Id()
	@Column(name = "id")
	private UUID id;

    @Column(name = "number")
    @CpfCnpjValidation
    private String number;

    @Column(name = "blocked")
    private Boolean blocked;
    
}
