rabbitmq:
	docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.11-management

rabbitmq-add:
	docker exec rabbitmq rabbitmqctl add_user jake secret

rabbitmq-tags:
	docker exec rabbitmq rabbitmqctl set_user_tags jake administrator

rabbitmq-guest:
	docker exec rabbitmq rabbitmqctl delete_user guest

rabbitmq-vhost:
	docker exec rabbitmq rabbitmqctl add_vhost customers

rabbitmq-perms:
	docker exec rabbitmq rabbitmqctl set_permissions -p customers jake ".*" ".*" ".*"

rabbitmq-admin-topic:
	docker exec rabbitmq rabbitmqadmin declare exchange --vhost=customers name=customer_events type=topic -u jake -p secret durable=true

rabbitmq-topic-perms:
	docker exec rabbitmq rabbitmqctl set_topic_permissions -p customers jake customer_events "^customers.*" "^customers.*"

rabbitmq-admin-delete:
	docker exec rabbitmq rabbitmqadmin delete exchange name=customer_events --vhost=customers -u jake -p secret

rabbitmq-admin-fanout:
	docker exec rabbitmq rabbitmqadmin declare exchange --vhost=customers name=customer_events type=fanout -u jake -p secret durable=true

rabbitmq-topic-perms-2:
	docker exec rabbitmq rabbitmqctl set_topic_permissions -p customers jake customer_events ".*" ".*"

rabbitmq-admin-direct:
	docker exec rabbitmq rabbitmqadmin declare exchange --vhost=customers name=customer_callbacks type=direct -u jake -p secret durable=true

rabbitmq-topic-perms-3:
	docker exec rabbitmq rabbitmqctl set_topic_permissions -p customers jake customer_callbacks ".*" ".*"

tls:
	git clone https://github.com/rabbitmq/tls-gen tls-gen && \
    cd tls-gen/basic && \
    make PASSWORD= && \
    make verify && \
    sudo chmod 644 tls-gen/basic/result/*

rabbitmq-remove:
	docker container rm -f rabbitmq

rabbitmq-tls:
	docker run -d \
	--name rabbitmq \
	-v "$(pwd)"/rabbitmq.conf:/etc/rabbitmq.conf:ro \
	-v "$(pwd)"/tls-gen/basic/result:/certs \
	-p 5671:5671 \
	-p 15672:15672 \
	rabbitmq:3.11-management

rabbitmq-restart:
	docker restart rabbitmq

rabbitmq-def:
	docker run -d \
	--name rabbitmq \
	-v "$(pwd)"/rabbitmq_definitions.json:/etc/rabbitmq/rabbitmq_definitions.json:ro \
	-v "$(pwd)"/rabbitmq.conf:/etc/rabbitmq/rabbitmq.conf:ro \
	-v "$(pwd)"/tls-gen/basic/result:/certs \
	-p 5671:5671 \
	-p 15672:15672 \
	rabbitmq:3.11-management