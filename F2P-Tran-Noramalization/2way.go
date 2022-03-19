Product Name: Bread
lot 1 : (3 SGTIN)


1) F2P-1 :{
	id:F2P-1
	orgId:1
	docType: 'Product'
	createdDate: "11/11/2021"
}

Commission
 LOT1:{
	_id: uuid30
	productId: F2P-1
	lotId: LOT1
	docType: "LOT"
	status: "Available"
	orgId:1
	CTE: "Commission"
	createdDate: 11/11/2021,
	}

LOT2:{
	_id: uuid31
	productId: F2P-1
	lotId: LOT2
	docType: "LOT"
	status: "Available"
	orgId:1
	CTE: "Commission"
	createdDate: 11/11/2021,
	}

LOT1: {
	id: LOT1,
	events: [
		uuid30,
		uuid1,
		uuid3,
		uuid5,
		uuid6
	]
}
LOT2: {
	id: LOT2,
	events: [
		uuid31,
		uuid2,
		uuid3,
		uuid5,
		uuid6
	]
}


SGTIN
 LOT1-SGTIN1: {
	_id: LOT1-SGTIN1
	data:["uuid30"]
	}

SGTIN
 LOT1-SGTIN2: {
	_id: LOT1-SGTIN2
	data:["uuid30"]
	}

SGTIN
 LOT1-SGTIN3: {
	_id:LOT1-SGTIN3
	data:["uuid30"]
	}

 LOT2-SGTIN1: {
_id:LOT2-SGTIN1
data:["uuid31"]
}


Observation:{
	_id: uuid1
	productId: F2P-1,
	lotId: LOT1,
	orgId:1
	docType: "EVENT"
	CTE: "Observation"
	createdDate: 11/11/2021,
	ObservationKDK: ""
}

Observation:{
	_id: uuid2
	productId: F2P-1,
	lotId: LOT2,
	orgId:1
	docType: "Observation"
	CTE: "Commission"
	createdDate: 11/11/2021,
	ObservationKDK: ""
}

At the stage of aggregation create two typ of tx need to create
	1) Aggregation
	2) SSCC
		Aggregaion: {
			productId:F2P-1,
			lotId: LOT1,
			orgId:1
			_id: uuid3
			docType: "Aggregaion"
			CTE: "Aggregaion"
			createdDate: 11/11/2021
			SSCC: 12345,
			InputKDE: [
				LOT1,
				LOT2
			]
		}

		SSCC:{
			id: 12345,
			lots: [uuid3],

		}
====================================

Manufacturer
Product: Bread
productId: F2P-2

Disaggregate: {
	productId:F2P-1,
	_id: uuid5
	lotId: LOT1,
	orgId:2
	SSCC: 12345
	docType: "Disaggregate"
	CTE: "Disaggregate"
	createdDate: 11/11/2021
	outputKDE: [LOT1, LOT2]
}




Transformtion: {
	_id: uuid6
	productId:F2P-2,
	orgId:2
	lotId: LOT3
	docType: "Transformtion"
	CTE: "Transformtion"
	createdDate: 11/11/2021
	input: [
			{
				productId:F2P-1,
				_id: uuid4
				lotId: LOT1,
			},
			{
				productId:F2P-1,
				_id: uuid5
				lotId: LOT2,
			}
		]
}

LOT3: {
	id: LOT3,
	events: [uuid6, uuid2, uuid33, uuid34]
}

SGTIN
transformation:{
	_id : LOT3-SGTIN1
	data: [uuid6]
}

SGTIN
transformation:{
	_id : LOT3-SGTIN2
	data: [uuid6]
}

Observation:{
	_id: uuid2
	productId: F2P-2,
	lotId: LOT3,
	orgId:2
	docType: "Observation"
	CTE: "Observation"
	createdDate: 11/11/2021,
	ObservationKDK: ""
}

Aggregaion: {
		productId:F2P-2,
		lotId: LOT3,
		orgId:2
		_id: uuid33
		docType: "Aggregaion"
		CTE: "Aggregaion"
		createdDate: 11/11/2021
		SSCC: 12346
		inputKDE: [LOT3]
	}
	SSCC:{
		id: 12346,
		lots: [uuid33],

		}
==========================================

Retailer:
	
Disaggregaion: {
		productId:F2P-2,
		lotId: LOT3,
		orgId:3
		_id: uuid34
		docType: "Aggregaion"
		CTE: "Aggregaion"
		createdDate: 11/11/2021
		SSCC: 12346
	}



=======================================

Traceability :--------

LOT3-SGTIN1: {
	uuid6:{
		input: {
				LOT1:[	uuid30,
						uuid1,
						uuid3,
						uuid5,
						uuid6
					  ]
				, 
				LOT2:[	uuid31,
						uuid2,
						uuid3,
						uuid5,
						uuid6]
			},
		output: {
				LOT3: [uuid6, uuid2, uuid33, uuid34]
			}
	}
}
LOT3-SGTIN1:{
	product Details: {id: F2P-2},

	Transformtion: {
	_id: LOT3
	productId:F2P-2,
	orgId:2
	docType: "Transformtion"
	CTE: "Transformtion"
	createdDate: 11/11/2021
	input: [
			{productId:F2P-1,
			_id: uuid4
			lotId: LOT1,
			},
			{productId:F2P-1,
			_id: uuid5
			lotId: LOT2,
			}
		]
	},
	Disaggregate: {
		productId:F2P-1,
		_id: uuid5
		lotId: LOT1,
		orgId:2
		SSCC: 12345
		docType: "Disaggregate"
		CTE: "Disaggregate"
		createdDate: 11/11/2021
	}

	Disaggregate: {
		productId:F2P-1,
		_id: uuid4
		SSCC: 12345
		lotId: LOT2,
		orgId:2
		docType: "Disaggregate"
		CTE: "Disaggregate"
		createdDate: 11/11/2021
	}

	Aggregaion: {
			productId:F2P-1,
			lotId: LOT1,
			orgId:1
			_id: uuid3
			docType: "Aggregaion"
			CTE: "Aggregaion"
			createdDate: 11/11/2021
			SSCC: 12345
		}

	Aggregaion: {
		productId:F2P-1,
		_id: uuid4
		lotId: LOT2,
		orgId:1
		docType: "Aggregaion"
		CTE: "Aggregaion"
		createdDate: 11/11/2021
		SSCC: 12345
	}

}
