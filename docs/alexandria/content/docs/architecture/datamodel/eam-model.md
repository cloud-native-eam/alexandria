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
|  belongs to	|  ABC	| ABO 	|  	|  	|
|  	|  	|  	|  	|  	|

## Business Capability (ABC)
|  Connection	|  from	| to |  Information attached	| Description|
|-	|-	|-	|-	|- |
| consists of 	| ABC 	|  ABC	|  	| An business capability can have multiple layers, the `consist of` explains this breakdown 	|
|  belongs to	|  ABC	| ABO 	|  	|  	|
|  	|  	|  	|  	|  	|

## Product (Application) (APA)
|  Connection	|  from	| to |  Information attached	| Description|
|-	|-	|-	|-	|- |
| fulfills 	| APA 	|  ABO	|  	|  	|
|  implements	|  AC	| APA 	|  	|  	|
| serves 	| ATS 	| APA 	|  	|  	|
| runs on 	| APA 	| AP 	|  	|  	|

## Component (AC)
|  Connection	|  from	| to |  Information attached	| Description|
|-	|-	|-	|-	|- |
| implements 	| AC 	|  APA	|  	|  	|
|  implements	|  AC	| ATS 	|  	|  	|
|  consists of	|  AC	|  AT	|  	|  	|

## (Technical) Service (ATS)
|  Connection	|  from	| to |  Information attached	| Description|
|-	|-	|-	|-	|- |
| serves 	| ATS 	|  APA	|  	|  	|
|  consits of	|  ATS	| AT 	|  	|  	|
|  runs on	|  ATS	|  AP	|  	|  	|
|  implements	|  AC	| ATS 	|  	|  	|
|  offers	|  AV	|  ATS	|  	|  	|

## Technology (AT)
|  Connection	|  from	| to |  Information attached	| Description|
|-	|-	|-	|-	|- |
| consists of 	| AP 	|  AT	|  	|  	|
|  consists of	|  ATS	| AT 	|  	|  	|
|  consists of	|  AC	|  AT	|  	|  	|
|  develop	|  AV	|  AT	|  	|  	|

## Platform (AP)
|  Connection	|  from	| to |  Information attached	| Description|
|-	|-	|-	|-	|- |
| consists of 	| AP 	|  AT	|  	|  	|
|  runs on	|  APA	| AP 	|  	|  	|
|  runs on	|  ATS	|  AP	|  	|  	|
|  operate	|  AV	|  AP	|  	|  	|

## Vendor (AV)
|  Connection	|  from	| to |  Information attached	| Description|
|-	|-	|-	|-	|- |
| offers 	| AV 	|  ATS	|  	|  	|
|  develop	|  AV	| AT 	|  	|  	|
|  operate	|  AV	|  AP	|  	|  	|