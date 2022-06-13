package com.github.m2tx.neowaytc.backend.controller;

import java.util.Optional;
import java.util.UUID;

import javax.validation.Valid;

import org.springframework.graphql.data.method.annotation.Argument;
import org.springframework.graphql.data.method.annotation.MutationMapping;
import org.springframework.graphql.data.method.annotation.QueryMapping;
import org.springframework.stereotype.Controller;

import com.github.m2tx.neowaytc.backend.exceptions.IdentificationNumberAlreadyExistsException;
import com.github.m2tx.neowaytc.backend.exceptions.IdentificationNumberUpdateException;
import com.github.m2tx.neowaytc.backend.model.IdentificationNumber;
import com.github.m2tx.neowaytc.backend.service.IdentificationNumberService;

import lombok.AllArgsConstructor;

@Controller
@AllArgsConstructor
public class IdentificationNumberGraphQlController {

    private final IdentificationNumberService service;

    @QueryMapping("allIdentificationNumber")
    public Iterable<IdentificationNumber> allIdentificationNumber(){
        return service.findAll();
    }

    @QueryMapping("getIdentificationNumberByID")
    public Optional<IdentificationNumber> getIdentificationNumberByID(@Argument("id") UUID id){
        return service.findById(id);
    }

    @MutationMapping
    public IdentificationNumber newIdentificationNumber(@Argument("identificationNumber") @Valid IdentificationNumber identificationNumber) throws IdentificationNumberAlreadyExistsException{
        return service.newIdentificationNumber(identificationNumber);
    }

    @MutationMapping
    public IdentificationNumber updateIdentificationNumber(@Argument("identificationNumber") @Valid IdentificationNumber identificationNumber) throws IdentificationNumberUpdateException{
        service.updateIdentificationNumber(identificationNumber);
        return identificationNumber;
    }
    
}
