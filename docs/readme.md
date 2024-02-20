## This context is collector-service

### Entities
```
customer
{
	id: uuid,
	name: string
}

contract
{
	id: uuid,
	id_customer: fk customer.id
}

charges
{
	id: uuid,
	reference: date,
	expiration_date: date,
	payment_date: date,
	status: bool, // 1 -> paid | 0 -> unpaid
	id_contract: fk contract.id
}
```

### Obs
```
-- Criar estrutura base com docker compose
-- Variaveis de ambiente registradas em um .env.example
-- Utilizar Apache Kafka para Topicos

Service: negligent-service
Database: postgres

Service: collector-service
Database: mongodb
```

### Queues:
```
	1 -> (collector manda para o negligent o range de datas): request-bills
	{
		initial_date: yyyy-mm-dd,
		end_date: yyyy-mm-dd
	}

	2 -> (negligent manda os dados de cobrança para o collector): response-bills
	{
		reference: string,
		contract_id: number,
		customer_name: string,
		expiration_date: yyyy-mm-dd
	}

	3 -> (collector retorna para negligent mensalidades para baixa): register-payment
	{
		reference: string,
		contract_id: number,
		payment_date: yyyy-mm-dd
	}
```
