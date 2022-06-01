package com.xws.agentska.repository;

import com.xws.agentska.model.JobPosition;
import com.xws.agentska.model.Role;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;

public interface JobPositionRepository extends JpaRepository<JobPosition,Long> {
    @Query("select jp from JobPosition jp where jp.name=?1")
    public JobPosition findByName(String name);
}
