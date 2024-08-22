part of 'insert_product_bloc.dart';

 class ProductInsertedEvent {
  ProductInsertedEvent();
}

class ProductInserted extends  ProductInsertedEvent {
  ProductEntity product;
  ProductInserted({required this.product});
}
class InsertInitial extends  ProductInsertedEvent {
  InsertInitial();
}
