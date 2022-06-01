package com.xws.agentska.service.interfaces;

import com.xws.agentska.model.ApiConnection;
import com.xws.agentska.model.Company;
import com.xws.agentska.model.JobOffer;
import com.xws.agentska.model.enumerations.Status;

public interface CompanyService extends Service<Company,Long> {
    public void updateCompanyStatus(long companyId, Status newStatus);
    public void addNewJobOffer(long companyId,JobOffer offer);
    public void addNewApiConnection(long companyId, ApiConnection connection);
}
