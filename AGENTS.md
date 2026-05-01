# Recipe Service Agents Guide

This service owns recipe-related domain behavior.

## Scope

- recipe entities, storage, and business rules
- recipe creation, updates, retrieval, and related transport contracts
- interactions with encryption, tags, community, or collection flows when they are recipe-driven

## Working Rules

- Keep recipe domain rules centralized here instead of scattering them into gateway or client-facing glue code.
- Be explicit about ownership when a feature touches tags, shopping lists, or encryption.
- Review migration impact carefully because recipe data is likely central and high-volume.
