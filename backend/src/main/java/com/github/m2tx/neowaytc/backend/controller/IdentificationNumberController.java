package com.github.m2tx.neowaytc.backend.controller;

import java.net.URI;
import java.util.UUID;

import javax.validation.Valid;

import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.servlet.support.ServletUriComponentsBuilder;

import com.github.m2tx.neowaytc.backend.exceptions.IdentificationNumberAlreadyExistsException;
import com.github.m2tx.neowaytc.backend.exceptions.IdentificationNumberUpdateException;
import com.github.m2tx.neowaytc.backend.model.IdentificationNumber;
import com.github.m2tx.neowaytc.backend.repository.IdentificationNumberSpecification;
import com.github.m2tx.neowaytc.backend.service.IdentificationNumberService;

import lombok.AllArgsConstructor;

@RestController()
@RequestMapping("/")
@AllArgsConstructor
public class IdentificationNumberController {

    private final IdentificationNumberService service;

	@GetMapping()
	public Iterable<IdentificationNumber> getAll(){
		return service.findAll();
	}

	@GetMapping("/{id}")
	public ResponseEntity<IdentificationNumber> get(@PathVariable UUID id){
		return ResponseEntity.of(service.findById(id));
	}

	@PostMapping("/query/")
    public ResponseEntity<Page<IdentificationNumber>> query(@RequestBody IdentificationNumberSpecification identificationNumber, Pageable pageable) {
		return ResponseEntity.ok(service.query(identificationNumber,pageable));
    }

    @PostMapping()
    public ResponseEntity<IdentificationNumber> newIdentificationNumber(@Valid @RequestBody IdentificationNumber identificationNumber) throws IdentificationNumberAlreadyExistsException {
        service.newIdentificationNumber(identificationNumber);
		URI location = ServletUriComponentsBuilder.fromCurrentRequest()
                .path("/{id}")
                .buildAndExpand(identificationNumber.getId())
                .toUri();
		return ResponseEntity.created(location).body(identificationNumber);
    }

	@PutMapping()
    public ResponseEntity<IdentificationNumber> updateIdentificationNumber(@RequestBody IdentificationNumber identificationNumber) throws IdentificationNumberUpdateException {
        service.updateIdentificationNumber(identificationNumber);
		return ResponseEntity.ok(identificationNumber);
    }

	@DeleteMapping("/{id}")
	public ResponseEntity<UUID> delete(@PathVariable UUID id){
		service.deleteById(id);
		return ResponseEntity.ok(id);
	}
    
}
