package com.github.m2tx.neowaytc.backend.controller;

import static org.hamcrest.CoreMatchers.any;
import static org.hamcrest.Matchers.hasSize;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.delete;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.post;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.put;
import static org.springframework.test.web.servlet.result.MockMvcResultHandlers.print;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.jsonPath;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

import java.util.UUID;

import org.junit.jupiter.api.Test;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.ValueSource;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.jdbc.AutoConfigureTestDatabase;
import org.springframework.boot.test.autoconfigure.jdbc.AutoConfigureTestDatabase.Replace;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.http.MediaType;
import org.springframework.test.annotation.DirtiesContext;
import org.springframework.test.annotation.DirtiesContext.ClassMode;
import org.springframework.test.context.ActiveProfiles;
import org.springframework.test.web.servlet.MockMvc;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.github.m2tx.neowaytc.backend.model.IdentificationNumber;

@SpringBootTest
@AutoConfigureMockMvc
@DirtiesContext(classMode = ClassMode.BEFORE_EACH_TEST_METHOD)
@AutoConfigureTestDatabase(replace = Replace.ANY)
@ActiveProfiles("test")
public class IdentificationNumberControllerTest {
    
	@Autowired
	private MockMvc mockMvc;

	@Autowired
	private ObjectMapper mapper;

    @ParameterizedTest
    @ValueSource(strings = {
        "39264468030",
        "10117286052",
		"02523798000181",
        "93566231000148",
        "89168173000127",
        "73953500000111"
    })
	public void addValidIdentificationNumber(String number) throws JsonProcessingException, Exception {
		this.mockMvc.perform(post("/api/identificationnumber/")
				.contentType(MediaType.APPLICATION_JSON)
				.content(mapper.writeValueAsString(IdentificationNumber.builder().number(number).build())))
		.andDo(print())
		.andExpect(status().isCreated())
		.andExpect(jsonPath("$.number").value(number))
		.andExpect(jsonPath("$.id").isNotEmpty());
	}

	@ParameterizedTest
    @ValueSource(strings = {
        "046.847.189-81",
		"89927530000194",
        "200B",
        "1234"
    })
	public void failedToAddInvalidIdentificationNumber(String number) throws JsonProcessingException, Exception {
		this.mockMvc.perform(post("/api/identificationnumber/")
				.contentType(MediaType.APPLICATION_JSON)
				.content(mapper.writeValueAsString(IdentificationNumber.builder().number(number).build())))
		.andDo(print())
		.andExpect(status().isBadRequest())
		.andExpect(jsonPath("$[0]").value("Invalid number: must be CPF or CNPJ"));
	}

	@Test
	public void queryIdentificationNumber() throws JsonProcessingException, Exception {
		this.mockMvc.perform(post("/api/identificationnumber/query/")
				.contentType(MediaType.APPLICATION_JSON)
				.content(mapper.writeValueAsString(IdentificationNumber.builder().blocked(true).build())))
		.andDo(print())
		.andExpect(status().isOk())
		.andExpect(jsonPath("$.content",hasSize(1)));
	}

	@ParameterizedTest
    @ValueSource(strings = {
        "789c728f-8fa2-494b-8db1-18808a5c61d8",
		"8ccf972c-6f24-4df3-ac65-b94853c10744",
        "35240f60-6a08-4774-becd-826bae221876"
    })
	public void findByIdIdentificationNumber(String id) throws JsonProcessingException, Exception {
		this.mockMvc.perform(get("/api/identificationnumber/"+id)
				.contentType(MediaType.APPLICATION_JSON))
		.andDo(print())
		.andExpect(status().isOk())
		.andExpect(jsonPath("$.id").value(id));
	}

	@ParameterizedTest
    @ValueSource(strings = {
        "123e4567-e89b-12d3-a456-426614174000"
    })
	public void notFoundFindByIdIdentificationNumber(String id) throws JsonProcessingException, Exception {
		this.mockMvc.perform(get("/api/identificationnumber/"+id)
				.contentType(MediaType.APPLICATION_JSON))
		.andDo(print())
		.andExpect(status().isNotFound());
	}

	@ParameterizedTest
    @ValueSource(strings = {
		"789c728f-8fa2-494b-8db1-18808a5c61d8",
		"8ccf972c-6f24-4df3-ac65-b94853c10744",
        "35240f60-6a08-4774-becd-826bae221876"
    })
	public void updateBlockedIdentificationNumber(String id) throws JsonProcessingException, Exception {
		this.mockMvc.perform(put("/api/identificationnumber/"+id)
				.contentType(MediaType.APPLICATION_JSON)
				.content(mapper.writeValueAsString(IdentificationNumber.builder().id(UUID.fromString(id)).blocked(true).build())))
		.andDo(print())
		.andExpect(status().isOk())
		.andExpect(jsonPath("$.id").value(id))
		.andExpect(jsonPath("$.blocked").value(true));
	}

	@Test
	public void notFoundUpdateBlockedIdentificationNumber() throws JsonProcessingException, Exception {
		String id = "123e4567-e89b-12d3-a456-426614174000";
		this.mockMvc.perform(put("/api/identificationnumber/"+id)
				.contentType(MediaType.APPLICATION_JSON)
				.content(mapper.writeValueAsString(IdentificationNumber.builder().id(UUID.fromString(id)).blocked(true).build())))
		.andDo(print())
		.andExpect(status().isBadRequest())
		.andExpect(jsonPath("$[0]").value("Failed to update Identification Number!"));
	}

	@ParameterizedTest
    @ValueSource(strings = {
		"789c728f-8fa2-494b-8db1-18808a5c61d8",
		"8ccf972c-6f24-4df3-ac65-b94853c10744",
        "35240f60-6a08-4774-becd-826bae221876"
    })
	public void deleteIdentificationNumber(String id) throws JsonProcessingException, Exception {
		this.mockMvc.perform(delete("/api/identificationnumber/"+id)
				.contentType(MediaType.APPLICATION_JSON))
		.andDo(print())
		.andExpect(status().isOk())
		.andExpect(jsonPath("$").value(id));
	}

}
