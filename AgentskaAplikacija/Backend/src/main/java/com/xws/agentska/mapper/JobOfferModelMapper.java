package com.xws.agentska.mapper;

import com.xws.agentska.dto.JobOfferDto;
import com.xws.agentska.model.JobOffer;
import org.springframework.stereotype.Component;

@Component("jobOfferModelMapper")
public class JobOfferModelMapper extends CustomModelMapperAbstract<JobOffer, JobOfferDto> {
    @Override
    public JobOfferDto convertToDto(JobOffer entity) {

        return modelMapper.map(entity,JobOfferDto.class);
    }

    @Override
    public JobOffer convertToEntity(JobOfferDto dto) {
        JobOffer offer=modelMapper.map(dto,JobOffer.class);
        return offer;
    }
}
