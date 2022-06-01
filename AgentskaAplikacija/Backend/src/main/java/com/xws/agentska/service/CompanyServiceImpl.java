package com.xws.agentska.service;

import com.xws.agentska.dto.JobOfferDislinktDto;
import com.xws.agentska.model.*;
import com.xws.agentska.model.enumerations.Status;
import com.xws.agentska.service.interfaces.CompanyService;
import com.xws.agentska.service.interfaces.JobPositionService;
import com.xws.agentska.service.interfaces.UsersService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.web.client.RestTemplateBuilder;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpHeaders;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;

import javax.transaction.Transactional;
import java.sql.Timestamp;
import java.time.Instant;

@Service
public class CompanyServiceImpl extends CustomGenericService<Company, Long> implements CompanyService {
    @Autowired
    private UsersService usersService;

    @Autowired
    private JobPositionService jobPositionService;
    @Autowired
    private RestTemplateBuilder restTemplateBuilder;

    @Override
    public void save(Company item) {
        item.setStatus(Status.PENDING);
        item.setOwner(usersService.findById(item.getOwner().getId()));
        super.save(item);
    }

    @Transactional
    @Override
    public void updateCompanyStatus(long companyId, Status newStatus) {
        Company company = this.findById(companyId);
        company.setStatus(newStatus);
        this.update(company);
        usersService.updateUserRole(company.getOwner().getId(), new Role(Role.COMPANY_OWNER_ROLE));
    }

    @Override
    public void addNewJobOffer(long companyId, JobOffer offer, boolean share) {
        Company company = this.findById(companyId);
        offer.setCreatedAt(Timestamp.from(Instant.now()));
        setJobPositionIfExistsInDb(offer);
        company.addNewJobOffer(offer);
        this.update(company);
        if (share) this.shareOnDislinkt(offer, company.getApiConnection());
    }

    private void shareOnDislinkt(JobOffer offer, ApiConnection connectionInfo) {
        RestTemplate restTemplate=restTemplateBuilder.build();
        HttpHeaders httpHeaders=new HttpHeaders();
        httpHeaders.setBearerAuth(connectionInfo.getApiKey());
        JobOfferDislinktDto bodyOffer=new JobOfferDislinktDto(offer);
        HttpEntity<JobOfferDislinktDto> entity = new HttpEntity<>(bodyOffer, httpHeaders);
        restTemplate.postForObject(connectionInfo.getApi(), entity, JobOfferDislinktDto.class);
    }

    private void setJobPositionIfExistsInDb(JobOffer offer) {
        JobPosition position = jobPositionService.findByName(offer.getPosition().getName());
        if (position != null) offer.setPosition(position);
    }

    @Override
    public void addNewApiConnection(long companyId, ApiConnection connection) {
        Company company = this.findById(companyId);
        company.setApiConnection(connection);
        this.update(company);
    }
}
