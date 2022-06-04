package com.xws.agentska.controller;

import com.xws.agentska.dto.JobOfferDto;
import com.xws.agentska.mapper.JobOfferModelMapper;
import com.xws.agentska.model.ApiConnection;
import com.xws.agentska.model.Company;
import com.xws.agentska.model.User;
import com.xws.agentska.model.enumerations.Status;
import com.xws.agentska.security.TokenBasedAuthentication;
import com.xws.agentska.service.interfaces.CompanyService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import static com.xws.agentska.controller.UserController.getLoggedinUserId;

@RestController
public class CompanyController {
    @Autowired
    private CompanyService service;

    @Autowired
    private JobOfferModelMapper jobOfferModelMapper;

    @GetMapping("/api/companies")
    public ResponseEntity<?> findAll() {
        return ResponseEntity.ok().body(service.findAll());
    }

    @PostMapping("/api/companies")
    public ResponseEntity<?> createNewCompanyRequest(@RequestBody Company company) {
        company.setOwner(new User(getLoggedinUserId()));
        service.save(company);
        return ResponseEntity.ok().build();
    }

    @GetMapping("/api/users/{id}/companies")
    public ResponseEntity<?> findCompanyByOwnerId(@PathVariable long id) {
        return ResponseEntity.ok().body(service.findByOwnerId(id));
    }

    @GetMapping("/api/companies/{id}")
    public ResponseEntity<?> findById(@PathVariable long id) {
        return ResponseEntity.ok().body(service.findById(id));
    }

    @PutMapping("/api/companies/{id}")
    public ResponseEntity<?> createResponseToNewCompanyRequest(@PathVariable long id) {
        service.updateCompanyStatus(id, Status.ADMIN_CONFIRMED);
        return ResponseEntity.ok().build();
    }

    @PostMapping("/api/companies/{id}/job-offers")
    public ResponseEntity<?> createNewJobOffer(@PathVariable long id, @RequestBody JobOfferDto dto) {
        service.addNewJobOffer(id, jobOfferModelMapper.convertToEntity(dto), dto.isShareOnDislinkt());
        return ResponseEntity.ok().build();
    }

    @PostMapping("/api/companies/{id}/api-connections")
    public ResponseEntity<?> addNewConnection(@PathVariable long id, @RequestBody ApiConnection conn) {
        service.addNewApiConnection(id, conn);
        return ResponseEntity.ok().build();
    }
}
