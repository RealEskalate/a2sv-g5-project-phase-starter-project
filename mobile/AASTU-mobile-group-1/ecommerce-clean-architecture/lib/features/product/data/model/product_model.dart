import '../../domain/entities/product.dart';

class ProductModel extends Productentity{
const ProductModel({
  required String id,
  required  String image,
  required String name,
  required String description,
  required double price,}
) : super(
  id: id, 
  image: image,
   name: name, 
   description: description,
   price: price
   );

   factory ProductModel.fromJson(Map<String,dynamic>json){ 

    return ProductModel(
      id: json['id'],
      image: json['imageUrl'],
      name: json['name'],
      description: json['description'],
      price:( json['price'] as num?)?.toDouble() ?? 45.0,        
   );

}

     Map<String,dynamic> toJson()=>{
      'id': id,
      'name': name,
      'description': description,
      'price': price,
      'imageUrl': image,
    };

    factory ProductModel.fromEntity(Productentity entity) {
    return ProductModel(
      id: entity.id,
      image: entity.image,
      name: entity.name,
      description: entity.description,
      price: entity.price,
    );
  }
    
}
    
