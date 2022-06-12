package com.github.m2tx.neowaytc.backend.exceptions;

public class IdentificationNumberUpdateException extends Exception{
    
    public IdentificationNumberUpdateException(){
        super("Failed to update Identification Number!");
    }

    public IdentificationNumberUpdateException(String message){
        super(message);
    }

}
