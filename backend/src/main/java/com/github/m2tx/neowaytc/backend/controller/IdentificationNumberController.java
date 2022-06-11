package com.github.m2tx.neowaytc.backend.controller;

import java.net.URI;
import java.util.List;
import java.util.UUID;

import javax.validation.Valid;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.servlet.support.ServletUriComponentsBuilder;

import com.github.m2tx.neowaytc.backend.exceptions.IdentificationNumberAlreadyExistsException;
import com.github.m2tx.neowaytc.backend.model.IdentificationNumber;
import com.github.m2tx.neowaytc.backend.service.IdentificationNumberService;

import lombok.AllArgsConstructor;

@RestController()
@RequestMapping("/")
@AllArgsConstructor
public class IdentificationNumberController {

    private final IdentificationNumberService service;

	@GetMapping()
	public List<IdentificationNumber> getAll(){
		return service.findAll();
	}

	@GetMapping("/{id}")
	public ResponseEntity<IdentificationNumber> get(@PathVariable UUID id){
		return ResponseEntity.of(service.findById(id));
	}

    @PostMapping()
    public ResponseEntity<IdentificationNumber> addDocument(@Valid @RequestBody IdentificationNumber identificationNumber) throws IdentificationNumberAlreadyExistsException {
        service.add(identificationNumber);
		URI location = ServletUriComponentsBuilder.fromCurrentRequest()
                .path("/{id}")
                .buildAndExpand(identificationNumber.getId())
                .toUri();
		return ResponseEntity.created(location).body(identificationNumber);
    }
    
}
