package com.xws.agentska.service;

import com.xws.agentska.model.JobPosition;
import com.xws.agentska.repository.JobPositionRepository;
import com.xws.agentska.service.interfaces.JobPositionService;
import org.springframework.stereotype.Service;

@Service
public class JobPositionServiceImpl extends CustomGenericService<JobPosition,Long> implements JobPositionService {
    @Override
    public JobPosition findByName(String name) {
        return ((JobPositionRepository)repository).findByName(name);
    }
}
