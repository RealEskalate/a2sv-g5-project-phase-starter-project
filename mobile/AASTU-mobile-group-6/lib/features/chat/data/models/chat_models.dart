import 'package:ecommerce_app_ca_tdd/features/chat/domain/entities/chat_entity.dart';
import 'package:ecommerce_app_ca_tdd/features/product/data/models/seller_model.dart';

class ChatModel{
  final String id;
  final String name;
  final String description;
  final num price;
  final String imagePath;
  final SellerModel sellerId;
  final String user1, user2;
  ChatModel({
    required this.user1,
    required this.user2,
    required this.id,
    required this.name,
    required this.description,
    required this.price,
    required this.imagePath,
    required this.sellerId
  }) : super();

  factory ChatModel.fromJson(Map<String, dynamic> json) {
    return ChatModel(
      id : json['id'],// TO be deleteed
      name: json['name'],
      description: json['description'],
      price: (json['price']).toDouble(),
      imagePath: json['imageUrl'] ?? '',
      sellerId: SellerModel.fromJson(json['sellerId']),
      user1: json['user1'],
      user2: json['user2']
    );
  }

  Map<String, dynamic> toJson() => {
        "name": name,
        "description": description,
        "price": price,
        
      };
}