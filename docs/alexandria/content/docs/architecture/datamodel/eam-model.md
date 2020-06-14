---
title: EAM Model
weight: 1
bookToc: true
---

# Reference Model
The target is that the reference model is flexible in its constitution, nevertheless we will start with an idiomic approach and just keep the data fields within the categories flexible.

## Abbreviation
To keep docs short we will use abbreviation for the categories starting with an "A" for Alexandria and then capitalized the first latter from the category.

# Enterprise Architecture Reference Model
![Hihlevel Arch](/Reference_Model.png)

## Business Organization (ABO)
|  Connection	|  from	| to |  Information attached	| Description|
|-	|-	|-	|-	|- |
| consists of 	| ABO 	|  ABO	|  	| An organization can have multiple layers, the `consist of` explains this breakdown 	|
|  served by	|  APA	| ABO 	|  	| Responsible BO for the Application/Product (not who use it!) |
|  belongs to	|  ABC	| ABO 	|  	| Which BC is done by which BO	|
|  contracts	|  ABO	| AV 	|  	| Close contract to deliver/serve/develop/operate a component/technology/platform 	|

## Business Capability (ABC)
|  Connection	|  from	| to |  Information attached	| Description|
|-	|-	|-	|-	|- |
| consists of 	| ABC 	|  ABC	|  	| An business capability can have multiple layers, the `consist of` explains this breakdown 	|
|  belongs to	|  ABC	| ABO 	|  	| Which BC is done by which BO 	|
|  fulfills	|  APA	| ABC 	|  	| BC is solved by application/product	|

## Product (Application) (APA)
|  Connection	|  from	| to |  Information attached	| Description|
|-	|-	|-	|-	|- |
| fulfills 	| APA 	|  ABO	|  	| BC is solved by application/product 	|
| serves 	| AC 	| APA 	|  	| One or multiple components serves or getting abstract as a Product/application 	|
| runs on 	| APA 	| AP 	|  	| A Product/Application must be hosted on some infrastructure/platform 	|
| served by 	| APA 	| ABO 	|  	|  Responsible BO for the Application/Product (not who use it!) 	|

## Component (AC)
|  Connection	|  from	| to |  Information attached	| Description|
|-	|-	|-	|-	|- |
| serves 	| AC 	|  APA	|  	| One or multiple components serves or getting abstract as a Product/application	|
|  consits of	|  AC	| AT 	|  	| Logically abstractes a technology to a usable service 	|
|  runs on	|  AC	|  AP	|  	|  A service needs to run on an internal/external platform or as a SaaS	|
|  offers	|  AV	|  AC	|  	| Enterprise services for a component like Docker EE 	|

## Technology (AT)
|  Connection	|  from	| to |  Information attached	| Description|
|-	|-	|-	|-	|- |
| consists of 	| AP 	|  AT	|  	| If it's a known/internal platform, the tech stack is available	|
|  consists of	|  AC	|  AT	|  	| Logically abstracted by a component	|
|  develop	|  AV	|  AT	|  	| Is developed by a vendor	|

## Platform (AP)
|  Connection	|  from	| to |  Information attached	| Description|
|-	|-	|-	|-	|- |
| consists of 	| AP 	|  AT	|  	| If it's a known/internal platform, the tech stack is available	|
|  runs on	|  APA	| AP 	|  	| A Product must be hosted on some infrastructure/platform	|
|  runs on	|  AC	|  AP	|  	| A service needs to run on an internal/external platform or as a SaaS	|
|  operate	|  AV	|  AP	|  	| Is offered by/responsible for the operation of a platform and bound to resposnibilities via an SLA 	|

## Vendor (AV)
|  Connection	|  from	| to |  Information attached	| Description|
|-	|-	|-	|-	|- |
| offers 	| AV 	|  AC	|  	| Enterprise services for a component like Docker EE	|
|  develop	|  AV	| AT 	|  	| Develops a specific technology	|
|  operate	|  AV	|  AP	|  	| Offers/is responsible for the operation of a platform and bound to resposnibilities via an SLA 	|
|  contracts	|  ABO	| AV 	|  	| Close contract to deliver/serve/develop/operate a component/technology/platform 	|