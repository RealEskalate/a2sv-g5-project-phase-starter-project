

part of 'product_bloc.dart';



abstract class ProductState extends Equatable {
  const ProductState();
  
  @override
  List<Object> get props => [];
}


class homeloading extends ProductState {
  homeloading();

}
class homeloaded extends ProductState {
  final List<ProductModel> product;
  homeloaded(this.product);
}
class homefailure extends ProductState {
  final String message;
  homefailure(this.message);
}

class adding extends ProductState {
  adding();

}
class added extends ProductState {
  added();
  }
class addfailure extends ProductState {
  final String message;
  addfailure(this.message);
}

class updating extends ProductState {
  updating();

}
class updated extends ProductState {
  updated();
}
class updatefailure extends ProductState {
  final String message;
  updatefailure(this.message);
}

class deleting extends ProductState {
  deleting();

}
class deleted extends ProductState {
  deleted();
  }
class deletefailure extends ProductState {
  final String message;
  deletefailure(this.message);
}
class getloading extends ProductState {
  getloading();

}
class getloaded extends ProductState {
  final ProductModel product;
  getloaded(this.product);
}
 class getfailure extends ProductState {
  final String message;
  getfailure(this.message);
 }