package com.xws.agentska.service;

import com.xws.agentska.model.JobOffer;
import com.xws.agentska.service.interfaces.JobOfferService;
import org.springframework.stereotype.Service;

@Service
public class JobOfferServiceImpl extends CustomGenericService<JobOffer,Long> implements JobOfferService {
}
