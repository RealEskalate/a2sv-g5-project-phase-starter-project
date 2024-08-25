import 'dart:ffi';
import 'dart:io';

import 'package:ecommerce_app_ca_tdd/features/product/data/models/seller_model.dart';
import 'package:ecommerce_app_ca_tdd/features/product/domain/entities/product_entity_local.dart';
import 'package:flutter/foundation.dart';
import 'package:ecommerce_app_ca_tdd/features/product/domain/entities/product_entity.dart';
import 'package:image_picker/image_picker.dart';

class ProductModel extends ProductEntity_local {
  const ProductModel({
    required String id,
      required String name, required String description, required num price, required String imagePath,required SellerModel sellerId,
      
    }) : super(id:id,name:  name, description: description, price: price, imagePath: imagePath, sellerId: sellerId);

    


  factory ProductModel.fromJson(Map<String, dynamic> json) {
    return ProductModel(
      id : json['id'],// TO be deleteed
      name: json['name'],
      description: json['description'],
      price: (json['price']).toDouble(),
      imagePath: json['imageUrl'] ?? '',
      sellerId: SellerModel.fromJson(json['seller']),
    );
  }

  // Map<String, dynamic> toJson() {
  //   return {
  //     'name': name,
  //     'price': price.toString(),
  //     'description': description,
  //     'imageUrl': imagePath.toString(),
  //   };
  // }
  Map<String, dynamic> toJson() => {
        "name": name,
        "description": description,
        "price": price,
        
      };
  
}
