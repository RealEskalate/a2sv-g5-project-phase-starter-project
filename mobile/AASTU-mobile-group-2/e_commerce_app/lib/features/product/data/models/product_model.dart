import 'package:e_commerce_app/features/product/domain/entities/product.dart';

import '../../../auth/data/models/user_model.dart';

class ProductModel extends ProductEntity {
  ProductModel(
      {required String description,
      required String id,
      required String imageUrl,
      required String name,
      required double price,
      required UserModel seller
      
      })
      : super(
            description: description,
            name: name,
            id: id,
            imageUrl: imageUrl,
            price: price,
            seller: seller
            );
  factory ProductModel.fromJson(Map<String, dynamic> json) => ProductModel(
        description: json["description"] ?? '',
        id: json["id"] ?? '',
        imageUrl: json["imageUrl"] ?? '',
        name: json["name"] ?? '',
        price: (json["price"] as num?)?.toDouble() ?? 0.0,
        seller: UserModel.fromJson2(json['seller']),
      );

  static List<ProductModel> fromJsonList(List<dynamic> jsonList) {
    return jsonList.map((json) => ProductModel.fromJson(json)).toList();
  }

  Map<String, dynamic> toJson() => {
        "name": name,
        "description": description,
        "price": price,
        
      };

   ProductEntity toEntity() => ProductEntity(
      description: description,
      id: id,
      imageUrl: imageUrl,
      name: name,
      price: price,
      seller: seller
      
      
      );
  // static List<ProductEntity> toEntityList(List<dynamic> productModel){
  //   return productModel.map((product) => product.)
  // }
 
}
