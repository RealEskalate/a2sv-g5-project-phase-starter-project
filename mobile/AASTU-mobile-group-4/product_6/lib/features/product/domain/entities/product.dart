import 'package:equatable/equatable.dart';
import '../../../auth/domain/entity/auth_entity.dart';

class Product extends Equatable {
  final String id;
  final String name;
  final String description;
  final String imageUrl;
  final double price;
  final UserEntity seller;


  const Product({
    required this.id,
    required this.name,
    required this.description,
    required this.imageUrl,
    required this.price,
    required this.seller,


  });

// for bloc state management
  Product copyWith({
    String? id,
    String? name,
    String? description,
    String? imageUrl,
    double? price,
    UserEntity? seller,
  }) {
    return Product(
      id: id ?? this.id,
      name: name ?? this.name,
      description: description ?? this.description,
      imageUrl: imageUrl ?? this.imageUrl,
      price: price ?? this.price, 
      seller: seller ?? this.seller,
      
    );
  }

//model part 
  factory Product.fromJson(Map<String, dynamic> json) {
    return Product(
      id: json['id'],
      name: json['name'],
      description: json['description'],
      imageUrl: json['imageUrl'],
      price: json['price'].toDouble(),
      seller:json['seller'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'name': name,
      'description': description,
      'imageUrl': imageUrl,
      'price': price,
    };
  }

  @override
  List<Object?> get props => [id, name, description, imageUrl, price];
}
