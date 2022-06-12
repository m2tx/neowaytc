package com.github.m2tx.neowaytc.backend.service;

import java.util.Optional;
import java.util.UUID;

import org.springframework.data.domain.Example;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.stereotype.Service;

import com.github.m2tx.neowaytc.backend.exceptions.IdentificationNumberAlreadyExistsException;
import com.github.m2tx.neowaytc.backend.exceptions.IdentificationNumberUpdateException;
import com.github.m2tx.neowaytc.backend.model.IdentificationNumber;
import com.github.m2tx.neowaytc.backend.repository.IdentificationNumberRepository;
import com.github.m2tx.neowaytc.backend.repository.IdentificationNumberSpecification;

import lombok.AllArgsConstructor;

@Service
@AllArgsConstructor
public class IdentificationNumberService {

	private final IdentificationNumberRepository repository;

	public Optional<IdentificationNumber> findById(UUID id) {
		return repository.findById(id);
	}

	public IdentificationNumber newIdentificationNumber(IdentificationNumber identificationNumber) throws IdentificationNumberAlreadyExistsException {
		if(repository.exists(Example.of(identificationNumber))) {
			throw new IdentificationNumberAlreadyExistsException();
		}
		identificationNumber.setId(UUID.randomUUID());
		identificationNumber.setBlocked(false);
		return repository.save(identificationNumber);
	}

	public void deleteById(UUID id) {
		repository.deleteById(id);
	}

	public Iterable<IdentificationNumber> findAll() {
		return repository.findAll();
	}

    public Long count(){
        return repository.count();
    }

    public Page<IdentificationNumber> query(IdentificationNumberSpecification spec, Pageable pageable) {
		return repository.findAll(spec,pageable);
    }

    public void updateIdentificationNumber(IdentificationNumber identificationNumber) throws IdentificationNumberUpdateException {
		Optional<IdentificationNumber> opt = repository.findById(identificationNumber.getId());
		if(opt.isPresent()){
			IdentificationNumber idn = opt.get();
			idn.setBlocked(identificationNumber.getBlocked());
			repository.save(idn);
		}else{
			throw new IdentificationNumberUpdateException();
		}
    }
}
