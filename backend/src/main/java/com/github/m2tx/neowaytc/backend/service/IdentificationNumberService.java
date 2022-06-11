package com.github.m2tx.neowaytc.backend.service;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

import org.springframework.data.domain.Example;
import org.springframework.stereotype.Service;

import com.github.m2tx.neowaytc.backend.exceptions.IdentificationNumberAlreadyExistsException;
import com.github.m2tx.neowaytc.backend.model.IdentificationNumber;
import com.github.m2tx.neowaytc.backend.repository.IdentificationNumberRepository;

import lombok.AllArgsConstructor;

@Service
@AllArgsConstructor
public class IdentificationNumberService {

	private final IdentificationNumberRepository repository;

	public Optional<IdentificationNumber> findById(UUID id) {
		return repository.findById(id);
	}

	public IdentificationNumber add(IdentificationNumber identificationNumber) throws IdentificationNumberAlreadyExistsException {
		if(repository.exists(Example.of(identificationNumber))) {
			throw new IdentificationNumberAlreadyExistsException();
		}
		identificationNumber.setId(UUID.randomUUID());
		return repository.save(identificationNumber);
	}

	public void deleteById(UUID id) {
		repository.deleteById(id);
	}

	public List<IdentificationNumber> findAll() {
		return repository.findAll();
	}

    public Long count(){
        return repository.count();
    }
}