import 'dart:io';

import 'package:equatable/equatable.dart';

class ProductEntity_local extends Equatable {
  final String id;
  final String name;
  final String description;
  final num price;
  final String imagePath;
  const ProductEntity_local(
      {
      required this.id,
      required this.name,
      required this.description,
      required this.price,
      required this.imagePath});

  @override
  List<Object?> get props => [id,name, description, price,imagePath];


  Map<String, dynamic> toJson() => {
        "name": name,
        "description": description,
        "price": price,
        
      };
}
// Compare this snippet from mobile/natnael_wondwoesn/lib/features/product/data/repositories/product_repository.dart:
// import 'package:dartz/dartz.dart';
// import 'package:flutter_application_5/core/error/failures.dart';
// import 'package:flutter_application_5/features/product/domain/entities/product_entity.dart';
//
// abstract class ProductRepository {
//   Future<Either<Failure, String>> addProduct(String name, String description, double price, String imagePath);
//   Future<Either<Failure, String>> deleteProduct(String id);
//   Future<Either<Failure, String>> updateProduct(String id, String name, String description, double price, String imagePath);
//   Future<Either<Failure, ProductEntity>> getProduct(String id);
// }
