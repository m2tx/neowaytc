package com.github.m2tx.neowaytc.backend.repository;

import java.util.LinkedList;
import java.util.List;

import javax.persistence.criteria.CriteriaBuilder;
import javax.persistence.criteria.CriteriaQuery;
import javax.persistence.criteria.Predicate;
import javax.persistence.criteria.Root;

import org.apache.commons.lang3.StringUtils;
import org.springframework.data.jpa.domain.Specification;

import com.github.m2tx.neowaytc.backend.model.IdentificationNumber;

public class IdentificationNumberSpecification implements Specification<IdentificationNumber> {

    private String number;
    private Boolean blocked;

    @Override
    public Predicate toPredicate(Root<IdentificationNumber> root, CriteriaQuery<?> query, CriteriaBuilder builder) {
        List<Predicate> predicates = new LinkedList<>();
        if(StringUtils.isNotBlank(number)){
            predicates.add(builder.equal(root.get("number"), number));
        }
        if(blocked!=null){
            predicates.add(builder.equal(root.get("blocked"), blocked));
        }
        return builder.and(predicates.toArray(new Predicate[predicates.size()]));
    }
    
}
