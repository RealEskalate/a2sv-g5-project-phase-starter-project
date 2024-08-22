import 'package:bloc/bloc.dart';
import 'package:equatable/equatable.dart';
import '../../../domain/entities/product.dart';

import '../../../domain/use_case/get_all_products.dart';

part 'home_page_event.dart';
part 'home_page_state.dart';


class HomePageBloc extends Bloc<HomePageEvent, HomePageState> {
  final GetAllProducts getAllProducts;

  HomePageBloc({required this.getAllProducts}) : super(HomePageInitialState()) {
    on<FetchAllProductsEvent>(_onFetchAllProducts);
    on<AddProductToHomePageEvent>(_onAddProductToHomePageEvent);
  }

  Future<void> _onFetchAllProducts(FetchAllProductsEvent event, Emitter<HomePageState> emit) async {
    emit(HomePageLoadingState());

    final result = await getAllProducts();
    result.fold(
      (failure) => emit(const HomePageErrorState('Failed to fetch products')),
      (products) => emit(HomePageLoadedState(products)),
    );
  }

  Future<void> _onAddProductToHomePageEvent(AddProductToHomePageEvent event, Emitter<HomePageState> emit) async {
    if (state is HomePageLoadedState) {
      final updatedProducts = List<Product>.from((state as HomePageLoadedState).products)
        ..add(event.product);
      emit(HomePageLoadedState(updatedProducts));
    }
  }

  
}

