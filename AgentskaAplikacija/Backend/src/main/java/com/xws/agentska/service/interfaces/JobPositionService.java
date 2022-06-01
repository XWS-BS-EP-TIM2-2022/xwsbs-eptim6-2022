package com.xws.agentska.service.interfaces;

import com.xws.agentska.model.JobPosition;

public interface JobPositionService extends Service<JobPosition,Long>{
    public JobPosition findByName(String name);
}
