import 'package:bloc/bloc.dart';
import 'package:ecommerce_app_ca_tdd/core/usecases/usecases.dart';
import 'package:ecommerce_app_ca_tdd/features/product/domain/usecases/get_all_usecase.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/search/search_event.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/search/search_state.dart';




class SearchBloc extends Bloc<SearchEvent, SearchState> {
  final GetAllUsecase getAllProduct;

  SearchBloc(this.getAllProduct) : super(SearchInitial()) {
    on<LoadAllProductEvent>(_onLoadAllProducts);
    on<SearchProductEvent>(_onSearchProducts);
  }

  Future<void> _onLoadAllProducts(LoadAllProductEvent event, Emitter<SearchState> emit) async {
    emit(LoadingState());
    final products = await getAllProduct(NoParams());
    products.fold(
      (failure) => emit(FailedState(failure.message)),
      (productList) => emit(LoadedState(productList)),
    );
  }

  void _onSearchProducts(SearchProductEvent event, Emitter<SearchState> emit) {
    if (state is LoadedState) {
      final currentState = state as LoadedState;
      final filteredProducts = currentState.data.where((product) =>
          product.name.toLowerCase().contains(event.query.toLowerCase())).toList();
      emit(LoadedState(filteredProducts));
    }
  }
}