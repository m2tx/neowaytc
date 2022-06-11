package com.github.m2tx.neowaytc.backend.graphql.query;

import java.util.List;
import java.util.UUID;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import com.github.m2tx.neowaytc.backend.model.IdentificationNumber;
import com.github.m2tx.neowaytc.backend.service.IdentificationNumberService;

import graphql.kickstart.tools.GraphQLQueryResolver;

@Component
public class IdentificationNumberQuery implements GraphQLQueryResolver{
    
    private IdentificationNumberService service;

    @Autowired
    public IdentificationNumberQuery(IdentificationNumberService service) {
        this.service = service;
    }

    public List<IdentificationNumber> getIdentificationNumbers() {
        return service.findAll();
    }

    public IdentificationNumber getIdentificationNumber(UUID id) {
        return service.findById(id).orElse(null);
    }

}
