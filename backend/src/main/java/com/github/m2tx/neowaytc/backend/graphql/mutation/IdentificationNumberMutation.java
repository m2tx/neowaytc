package com.github.m2tx.neowaytc.backend.graphql.mutation;

import javax.validation.Valid;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import com.github.m2tx.neowaytc.backend.exceptions.IdentificationNumberAlreadyExistsException;
import com.github.m2tx.neowaytc.backend.model.IdentificationNumber;
import com.github.m2tx.neowaytc.backend.service.IdentificationNumberService;

import graphql.kickstart.tools.GraphQLMutationResolver;

@Component
public class IdentificationNumberMutation implements GraphQLMutationResolver{
    
    private IdentificationNumberService service;

    @Autowired
    public IdentificationNumberMutation(IdentificationNumberService service) {
        this.service = service;
    }

    public IdentificationNumber addIdentificationNumber(@Valid IdentificationNumber identificationNumber) throws IdentificationNumberAlreadyExistsException {
        return service.add(identificationNumber);
    }

}
