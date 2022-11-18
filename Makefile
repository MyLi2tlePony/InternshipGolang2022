up:
	cd deployments && docker-compose -f docker-compose.yaml build && docker-compose -f docker-compose.yaml up

integration-tests:
	cd deployments && docker-compose -f docker-compose.test.yaml build && docker-compose -f docker-compose.test.yaml up

.PHONY: up integration-tests