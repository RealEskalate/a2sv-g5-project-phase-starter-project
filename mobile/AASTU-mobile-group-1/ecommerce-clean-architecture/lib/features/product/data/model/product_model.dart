import '../../../auth/data/model/user_model.dart';
import '../../domain/entities/product.dart';

class ProductModel extends Productentity{
 ProductModel({
  required String id,
  required  String image,
  required String name,
  required String description,
  required double price,
  required UserModel seller,
  }
) : super(
  id: id, 
  image: image,
   name: name, 
   description: description,
   price: price,
    seller:UserModel.fromEntity(seller),
   );

   factory ProductModel.fromJson(Map<String,dynamic>json){ 

    return ProductModel(
      id: json['id'],
      image: json['imageUrl'],
      name: json['name'],
      description: json['description'],
      price:( json['price'] as num?)?.toDouble() ?? 45.0,   
      seller:  UserModel.forSeller(json['seller'])  
   );

}

     Map<String,dynamic> toJson()=>{
      'id': id,
      'name': name,
      'description': description,
      'price': price,
      'imageUrl': image,
      'seller': UserModel.fromEntity(seller).toJson(),
    };

    factory ProductModel.fromEntity(Productentity entity) {
    return ProductModel(
      id: entity.id,
      image: entity.image,
      name: entity.name,
      description: entity.description,
      price: entity.price,
      seller: UserModel.fromEntity(entity.seller),
    );
  }
    
}
    
